import React from 'react';
import ComponentCreator from '@docusaurus/ComponentCreator';

export default [
  {
    path: '/__docusaurus/debug',
    component: ComponentCreator('/__docusaurus/debug', '5ff'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/config',
    component: ComponentCreator('/__docusaurus/debug/config', '5ba'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/content',
    component: ComponentCreator('/__docusaurus/debug/content', 'a2b'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/globalData',
    component: ComponentCreator('/__docusaurus/debug/globalData', 'c3c'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/metadata',
    component: ComponentCreator('/__docusaurus/debug/metadata', '156'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/registry',
    component: ComponentCreator('/__docusaurus/debug/registry', '88c'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/routes',
    component: ComponentCreator('/__docusaurus/debug/routes', '000'),
    exact: true
  },
  {
    path: '/blog',
    component: ComponentCreator('/blog', 'b2f'),
    exact: true
  },
  {
    path: '/blog/archive',
    component: ComponentCreator('/blog/archive', '182'),
    exact: true
  },
  {
    path: '/blog/authors',
    component: ComponentCreator('/blog/authors', '0b7'),
    exact: true
  },
  {
    path: '/blog/authors/all-sebastien-lorber-articles',
    component: ComponentCreator('/blog/authors/all-sebastien-lorber-articles', '4a1'),
    exact: true
  },
  {
    path: '/blog/authors/yangshun',
    component: ComponentCreator('/blog/authors/yangshun', 'a68'),
    exact: true
  },
  {
    path: '/blog/first-blog-post',
    component: ComponentCreator('/blog/first-blog-post', '89a'),
    exact: true
  },
  {
    path: '/blog/long-blog-post',
    component: ComponentCreator('/blog/long-blog-post', '9ad'),
    exact: true
  },
  {
    path: '/blog/mdx-blog-post',
    component: ComponentCreator('/blog/mdx-blog-post', 'e9f'),
    exact: true
  },
  {
    path: '/blog/tags',
    component: ComponentCreator('/blog/tags', '287'),
    exact: true
  },
  {
    path: '/blog/tags/docusaurus',
    component: ComponentCreator('/blog/tags/docusaurus', '704'),
    exact: true
  },
  {
    path: '/blog/tags/facebook',
    component: ComponentCreator('/blog/tags/facebook', '858'),
    exact: true
  },
  {
    path: '/blog/tags/hello',
    component: ComponentCreator('/blog/tags/hello', '299'),
    exact: true
  },
  {
    path: '/blog/tags/hola',
    component: ComponentCreator('/blog/tags/hola', '00d'),
    exact: true
  },
  {
    path: '/blog/welcome',
    component: ComponentCreator('/blog/welcome', 'd2b'),
    exact: true
  },
  {
    path: '/markdown-page',
    component: ComponentCreator('/markdown-page', '3d7'),
    exact: true
  },
  {
    path: '/docs',
    component: ComponentCreator('/docs', 'e65'),
    routes: [
      {
        path: '/docs',
        component: ComponentCreator('/docs', '692'),
        routes: [
          {
            path: '/docs',
            component: ComponentCreator('/docs', 'eb2'),
            routes: [
              {
                path: '/docs/basics/add-new-token',
                component: ComponentCreator('/docs/basics/add-new-token', '02f'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/basics/basic-ops',
                component: ComponentCreator('/docs/basics/basic-ops', '503'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/basics/bridging/boba-token-bridge',
                component: ComponentCreator('/docs/basics/bridging/boba-token-bridge', 'c1f'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/basics/bridging/index',
                component: ComponentCreator('/docs/basics/bridging/index', '7bb'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/basics/bridging/light-bridge',
                component: ComponentCreator('/docs/basics/bridging/light-bridge', '57d'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/basics/bridging/standard-bridge',
                component: ComponentCreator('/docs/basics/bridging/standard-bridge', '4d4'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/basics/deploy-contract',
                component: ComponentCreator('/docs/basics/deploy-contract', '101'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/basics/index',
                component: ComponentCreator('/docs/basics/index', '45c'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/basics/json-rpc-api',
                component: ComponentCreator('/docs/basics/json-rpc-api', '924'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/basics/verify-sc',
                component: ComponentCreator('/docs/basics/verify-sc', 'e8a'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/faucets',
                component: ComponentCreator('/docs/faucets', 'f81'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/hc/addsub/',
                component: ComponentCreator('/docs/hc/addsub/', '096'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/hc/addsub/write-contract',
                component: ComponentCreator('/docs/hc/addsub/write-contract', '1a7'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/hc/addsub/write-handler',
                component: ComponentCreator('/docs/hc/addsub/write-handler', 'c9c'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/hc/deploy-contract',
                component: ComponentCreator('/docs/hc/deploy-contract', '634'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/hc/index',
                component: ComponentCreator('/docs/hc/index', 'a1d'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/hc/price-fetch',
                component: ComponentCreator('/docs/hc/price-fetch', 'f4c'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/hc/server-setup',
                component: ComponentCreator('/docs/hc/server-setup', '2ca'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/index',
                component: ComponentCreator('/docs/index', '2f2'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/node-operators/index',
                component: ComponentCreator('/docs/node-operators/index', 'a81'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/node-operators/network-upgrades',
                component: ComponentCreator('/docs/node-operators/network-upgrades', '566'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/node-operators/run-node-docker',
                component: ComponentCreator('/docs/node-operators/run-node-docker', '9f1'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/node-operators/run-node-source',
                component: ComponentCreator('/docs/node-operators/run-node-source', 'ebb'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/node-operators/snapshot-downloads',
                component: ComponentCreator('/docs/node-operators/snapshot-downloads', '708'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/docs/node-operators/software-release',
                component: ComponentCreator('/docs/node-operators/software-release', '93c'),
                exact: true,
                sidebar: "tutorialSidebar"
              }
            ]
          }
        ]
      }
    ]
  },
  {
    path: '/',
    component: ComponentCreator('/', '2e1'),
    exact: true
  },
  {
    path: '*',
    component: ComponentCreator('*'),
  },
];
