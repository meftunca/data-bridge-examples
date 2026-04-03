import type { Theme } from "vitepress";
import DefaultTheme from "vitepress/theme";

import { theme as OpenApiTheme } from "vitepress-openapi/client";
import "vitepress-openapi/dist/style.css";

import { enhanceAppWithTabs } from "vitepress-plugin-tabs/client";

export default {
  extends: DefaultTheme,
  enhanceApp({ app }) {
    OpenApiTheme.enhanceApp({ app });
    enhanceAppWithTabs(app);
  },
} satisfies Theme;
