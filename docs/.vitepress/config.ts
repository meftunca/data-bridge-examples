import { defineConfig } from "vitepress";
import sidebar from "../sidebar.json";

export default defineConfig({
  title: "Innovation Hub API",
  description:
    "DataBridge V2 ile otomatik üretilmiş API dokümantasyonu — 5 schema, 54 tablo, 14 RPC",
  base: "/data-bridge-examples/",

  head: [
    ["link", { rel: "icon", type: "image/svg+xml", href: "/logo.svg" }],
  ],

  themeConfig: {
    logo: "/logo.svg",
    siteTitle: "Innovation Hub",

    nav: [
      { text: "Ana Sayfa", link: "/" },
      { text: "API Reference", link: "/API_REFERENCE" },
      {
        text: "Schemas",
        items: [
          { text: "IAM", link: "/iam/" },
          { text: "Catalog", link: "/catalog/" },
          { text: "Orders", link: "/orders/" },
          { text: "Logistics", link: "/logistics/" },
          { text: "Analytics", link: "/analytics/" },
        ],
      },
      {
        text: "Swagger",
        link: "/swagger",
      },
    ],

    sidebar: transformSidebar(sidebar),

    socialLinks: [
      {
        icon: "github",
        link: "https://github.com/meftunca/data-bridge-examples",
      },
    ],

    search: {
      provider: "local",
    },

    editLink: {
      pattern:
        "https://github.com/meftunca/data-bridge-examples/edit/main/docs/:path",
      text: "Bu sayfayı düzenle",
    },

    footer: {
      message: "DataBridge V2 ile otomatik üretilmiştir",
      copyright: "© 2026 Maple Technologies",
    },

    outline: {
      level: [2, 3],
      label: "İçindekiler",
    },
  },

  markdown: {
    lineNumbers: true,
    image: { lazyLoading: true },
  },

  lastUpdated: true,
  cleanUrls: true,
  ignoreDeadLinks: true,
});

/**
 * sidebar.json /bg-docs/ prefix'li link'ler içeriyor.
 * VitePress base zaten /data-bridge-examples/ olduğundan,
 * link'lerdeki /bg-docs/ prefix'ini kaldırıyoruz.
 */
function transformSidebar(raw: Record<string, any[]>): Record<string, any[]> {
  const items = raw["/bg-docs/"] || [];
  return { "/": transformItems(items) };
}

function transformItems(items: any[]): any[] {
  return items.map((item) => {
    const result: any = { text: item.text };
    if (item.link) {
      result.link = item.link.replace(/^\/bg-docs/, "");
    }
    if (item.collapsed !== undefined) {
      result.collapsed = item.collapsed;
    }
    if (item.items) {
      result.items = transformItems(item.items);
    }
    return result;
  });
}
