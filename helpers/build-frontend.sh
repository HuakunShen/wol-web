# for building dist, executed with a docker-compose file
npm install -g @vue/cli
cd /wol-web/frontend
rm -rf dist node_modules
npm install
npm run build
