import { defineConfig } from "vitepress";
import { withMermaid } from "vitepress-plugin-mermaid";
import { tabsMarkdownPlugin } from "vitepress-plugin-tabs";
import sidebar from "../sidebar.json";

export default withMermaid(
  defineConfig({
    title: "Innovation Hub API",
    description:
      "Auto-generated API documentation with DataBridge V2 — 5 schemas, 54 tables, 14 RPC",
    base: "/data-bridge-examples/",

    head: [["link", { rel: "icon", type: "image/svg+xml", href: "/logo.svg" }]],

    themeConfig: {
      logo: "/logo.svg",
      siteTitle: "Innovation Hub",

      nav: [
        { text: "Home", link: "/" },
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
        {
          text: "Presentation",
          link: "/presentation/",
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
        text: "Edit this page",
      },

      footer: {
        message: "Auto-generated with DataBridge V2",
        copyright: "© 2026 Maple Technologies",
      },

      outline: {
        level: [2, 3],
        label: "On this page",
      },
    },

    markdown: {
      lineNumbers: true,
      image: { lazyLoading: true },
      config(md) {
        md.use(tabsMarkdownPlugin);
      },
    },

    lastUpdated: true,
    cleanUrls: true,
    ignoreDeadLinks: true,

    mermaid: {
      theme: "dark",
    },
  }),
);

/**
 * sidebar.json contains links with /bg-docs/ prefix.
 * Since VitePress base is already /data-bridge-examples/,
 * we strip the /bg-docs/ prefix from links.
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
