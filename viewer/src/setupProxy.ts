import { createProxyMiddleware } from 'http-proxy-middleware';
module.exports = function (app: any) {
  app.use(
    '/api/solver',
    createProxyMiddleware({
      target: 'http://solver:8080',
      changeOrigin: true,
    })
  );
};
