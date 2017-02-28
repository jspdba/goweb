#!/bin/bash
tar -xzvf web.tar.gz -C web
cd web
chmod +x web
exec ../restart.sh
