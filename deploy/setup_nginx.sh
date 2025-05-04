
sudo apt install nginx certbot python3-certbot-nginx
echo '
server {
  listen [::]:443 ssl; # managed by Certbot
  listen 443 ssl; # managed by Certbot
  ssl_certificate /etc/letsencrypt/live/kostjc.benalu.dev/fullchain.pem; # managed by Certbot
  ssl_certificate_key /etc/letsencrypt/live/kostjc.benalu.dev/privkey.pem; # managed by Certbot
  include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
  ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot


  gzip on;

  gzip_types
    text/css
    text/plain
    text/javascript
    application/javascript
    application/json
    application/x-javascript
    application/xml
    application/xml+rss
    application/xhtml+xml
    application/x-font-ttf
    application/x-font-opentype
    application/vnd.ms-fontobject
    image/svg+xml
    image/x-icon
    application/rss+xml;
  gzip_proxied    no-cache no-store private expired auth;
  gzip_min_length 1024;
  gzip_comp_level 2;
  gzip_buffers 32 8k;

  server_name kostjc.benalu.dev;

  location / {
    root /home/kostjc/svelte;
    try_files $uri @apiproxy;
  }

  location @apiproxy {
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_pass http://localhost:1236;
  }
}
' | sudo tee /etc/nginx/sites-available/kostjc.benalu.dev
sudo ln -s /etc/nginx/sites-available/kostjc.benalu.dev /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
sudo certbot --nginx -d kostjc.benalu.dev -d kostjc.benalu.dev
sudo systemctl enable certbot.timer

# note if error, comment out first the nginx ssl_* part