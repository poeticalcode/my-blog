git pull && yarn install && rm -r dist && yarn run prod && rm -r /opt/nginx/html/self_blog && cp -r dist/ /opt/nginx/html/self_blog