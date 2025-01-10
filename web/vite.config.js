import { defineConfig } from 'vite';
import { fileURLToPath, URL } from 'node:url';
import { resolve } from 'node:path';

const specificViewPath = 'templates';

export default defineConfig({
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  server: {
    hmr: true,
  },
  base: '/',
  build: {
    rollupOptions: {
      input: {
        main: resolve('./src/main.js'),
      },
      output: {
        dir: './static',
        entryFileNames: `js/[name].js`,
        assetFileNames: (assetInfo) => {
          if (assetInfo.name?.endsWith('.css')) {
            return `css/[name].[ext]`;
          }
          if (
            assetInfo.name?.endsWith('.woff') ||
            assetInfo.name?.endsWith('.woff2')
          ) {
            return 'fonts/[name].[ext]';
          }
          return `${specificViewPath}/[name].[ext]`;
        },
      },
    },
  },
});
