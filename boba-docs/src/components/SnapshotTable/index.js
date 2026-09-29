import React from 'react';
import {usePluginData} from '@docusaurus/useGlobalData';

/**
 * Renders the current reth snapshot for one chain, from data fetched at build
 * time by the snapshot-manifests plugin.
 *
 * Only the reth rows are generated. The Geth, Erigon and Legacy tables on this
 * page are hand-maintained historical artifacts with no publisher behind them,
 * so they stay as static markdown.
 */

function formatSize(bytes) {
  const gib = bytes / 1024 ** 3;
  return gib >= 10 ? `${gib.toFixed(0)} GiB` : `${gib.toFixed(1)} GiB`;
}

function formatDate(iso) {
  // Deliberately not toLocaleDateString: this runs during a static build, so
  // the output would silently depend on the build machine's locale.
  return iso.slice(0, 10);
}

export default function SnapshotTable({chain}) {
  const data = usePluginData('snapshot-manifests');
  const manifest = data && data.snapshots && data.snapshots[chain];

  // Fallback when the plugin omitted this chain. Says plainly that the table
  // could not be generated and points at the machine-readable source, rather
  // than rendering an empty table that looks like "no snapshots exist".
  if (!manifest) {
    const base = (data && data.baseUrl) || 'https://snapshots.boba.network';
    return (
      <admonition type="caution">
        <p>
          The current snapshot listing could not be generated for this build. The
          machine-readable manifest is always available at{' '}
          <a href={`${base}/${chain}/latest.json`}>
            {base}/{chain}/latest.json
          </a>
          .
        </p>
      </admonition>
    );
  }

  const {
    url,
    sha256,
    block_number: blockNumber,
    block_timestamp_utc: blockTime,
    size_compressed_bytes: size,
    includes_proofsdb: proofs,
    proofs_history_window_blocks: window,
  } = manifest;

  return (
    <>
      <table>
        <thead>
          <tr>
            <th>Client</th>
            <th>Snapshot Date</th>
            <th>Block</th>
            <th>Size</th>
            <th>Download Link</th>
            <th>Sha256sum</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>Reth</td>
            <td>{formatDate(blockTime)}</td>
            <td>{blockNumber.toLocaleString('en-US')}</td>
            <td>{formatSize(size)}</td>
            <td>
              <a href={url}>Link</a>
            </td>
            <td>
              <code>{sha256}</code>
            </td>
          </tr>
        </tbody>
      </table>
      <p>
        <small>
          Published weekly. This URL is immutable — it always serves the snapshot
          taken at block {blockNumber.toLocaleString('en-US')}, so a download can
          be resumed safely.{' '}
          {proofs && window ? (
            <>
              The archive includes a proofs database covering the{' '}
              {window.toLocaleString('en-US')} blocks (about{' '}
              {Math.round((window * 2) / 86400)} days) before that block;{' '}
              <code>eth_getProof</code> for earlier blocks will not be served
              until the node backfills them.
            </>
          ) : null}
        </small>
      </p>
    </>
  );
}
