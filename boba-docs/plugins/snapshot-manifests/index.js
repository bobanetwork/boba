/**
 * Fetches published reth snapshot manifests at BUILD time and exposes them as
 * plugin global data, so the snapshot downloads page renders real, current
 * figures as static HTML.
 *
 * WHY BUILD TIME AND NOT CLIENT SIDE: the table is reference material people
 * find by searching, so it should be in the HTML rather than assembled by JS
 * after load. It also means a reader never sees an empty table because an
 * endpoint was briefly unreachable.
 *
 * WHY NO CREDENTIALS: these are public, anonymous URLs. Nothing here needs a
 * token, which is what keeps this repo free of any credential worth stealing.
 * Freshness comes from the snapshot publisher POSTing a Vercel deploy hook
 * after a successful publish, so a rebuild happens when there is something new
 * to show rather than on a timer.
 *
 * FAILURE POLICY: a build must never fail because a snapshot endpoint is down.
 * Unreachable or invalid manifests are dropped with a loud warning and the page
 * falls back to its static text. Publishing a docs site with one table missing
 * is recoverable; blocking every docs deploy on R2 availability is not.
 */

const DEFAULT_CHAINS = [
  { id: 'boba-mainnet', label: 'BOBA Mainnet' },
  { id: 'boba-sepolia', label: 'BOBA Sepolia' },
];

const SHA256_RE = /^[0-9a-f]{64}$/;

/**
 * Reject anything we would be embarrassed to publish. The page instructs
 * operators to verify with `sha256sum`, so a row without a well-formed digest
 * is worse than no row: it invites them to trust an unverifiable download.
 *
 * Staleness is checked too. A publisher that silently stopped running is the
 * failure this whole pipeline exists to prevent -- the page it replaces sat at
 * a four-month-old snapshot because nobody noticed.
 */
function validate(manifest, chainId, maxAgeDays, warn) {
  const problems = [];

  if (!manifest || typeof manifest !== 'object') {
    problems.push('not a JSON object');
    return { ok: false, problems };
  }
  if (!SHA256_RE.test(String(manifest.sha256 || ''))) {
    problems.push(`sha256 is not 64 hex characters (got ${JSON.stringify(manifest.sha256)})`);
  }
  if (!Number.isInteger(manifest.block_number) || manifest.block_number <= 0) {
    problems.push(`block_number is not a positive integer (got ${JSON.stringify(manifest.block_number)})`);
  }
  if (!Number.isInteger(manifest.size_compressed_bytes) || manifest.size_compressed_bytes <= 0) {
    problems.push('size_compressed_bytes is not a positive integer');
  }
  if (typeof manifest.url !== 'string' || !manifest.url.startsWith('https://')) {
    problems.push('url is missing or not https');
  }

  const created = Date.parse(manifest.created_utc);
  if (Number.isNaN(created)) {
    problems.push(`created_utc is unparseable (got ${JSON.stringify(manifest.created_utc)})`);
  } else {
    const ageDays = (Date.now() - created) / 86400000;
    if (ageDays > maxAgeDays) {
      problems.push(
        `stale: published ${ageDays.toFixed(1)} days ago, limit is ${maxAgeDays}. ` +
          'The publisher may have stopped running.'
      );
    }
  }

  // The chain field should match the prefix it was fetched from. A mismatch
  // means the publisher is writing under the wrong key or the manifest was
  // built by an older pipeline -- worth surfacing, but not worth suppressing
  // an otherwise valid row over.
  if (manifest.chain && manifest.chain !== chainId) {
    warn(`manifest for ${chainId} reports chain "${manifest.chain}"`);
  }

  return { ok: problems.length === 0, problems };
}

module.exports = function snapshotManifestsPlugin(context, options) {
  // Nullish coalescing, not ||: the numeric options have meaningful zero
  // values (maxAgeDays: 0 means "reject anything not published this instant",
  // which is how the staleness check gets tested) and || would silently
  // substitute the default for them.
  const opts = options || {};
  const baseUrl = opts.baseUrl ?? 'https://snapshots.boba.network';
  const chains = opts.chains ?? DEFAULT_CHAINS;
  const maxAgeDays = opts.maxAgeDays ?? 21;
  const timeoutMs = opts.timeoutMs ?? 15000;

  return {
    name: 'snapshot-manifests',

    async loadContent() {
      const snapshots = {};

      await Promise.all(
        chains.map(async ({ id, label }) => {
          const url = `${baseUrl}/${id}/latest.json`;
          const warn = (msg) => console.warn(`[snapshot-manifests] ${msg}`);

          let manifest;
          try {
            const res = await fetch(url, {
              signal: AbortSignal.timeout(timeoutMs),
              headers: { accept: 'application/json' },
            });
            if (!res.ok) {
              warn(`${url} returned HTTP ${res.status}; omitting ${id}`);
              return;
            }
            manifest = await res.json();
          } catch (err) {
            warn(`${url} could not be fetched (${err.message}); omitting ${id}`);
            return;
          }

          const { ok, problems } = validate(manifest, id, maxAgeDays, warn);
          if (!ok) {
            warn(`${url} failed validation; omitting ${id}:`);
            problems.forEach((p) => warn(`    - ${p}`));
            return;
          }

          snapshots[id] = { ...manifest, label };
        })
      );

      if (Object.keys(snapshots).length === 0) {
        console.warn(
          '[snapshot-manifests] no usable manifests; the snapshot table will render its fallback'
        );
      }
      return { snapshots, baseUrl };
    },

    async contentLoaded({ content, actions }) {
      actions.setGlobalData(content);
    },
  };
};
