import DefaultTheme from "vitepress/theme";
import type { Theme } from "vitepress";

import { theme as OpenApiTheme } from "vitepress-openapi/client";
import "vitepress-openapi/dist/style.css";

export default {
  extends: DefaultTheme,
  enhanceApp({ app }) {
    OpenApiTheme.enhanceApp({ app });
  },
} satisfies Theme;
