#!/bin/bash

# Replace environment variables in the main.js file
JS_BUNDLE_FILE=$(find /usr/share/nginx/html/assets -name "*.js" -type f)

# Replace environment variables
if [ -n "$VITE_API_URL" ]; then
    sed -i "s|VITE_API_URL_PLACEHOLDER|$VITE_API_URL|g" $JS_BUNDLE_FILE
fi 