---
sort: 100
---
# 参考文档

[https://rundocs.io/](https://rundocs.io/)

# 我的配置方式
在仓库根目录添加以下四个文件：

1. Gemfile

```
source "https://gems.ruby-china.com"
gem "jekyll-rtd-theme"

gem "github-pages", group: :jekyll_plugins
```


2. Makefile

```
DEBUG=JEKYLL_GITHUB_TOKEN=blank PAGES_API_URL=http://0.0.0.0

default:
	@gem install jekyll bundler && bundle install

update:
	@bundle update

clean:
	@bundle exec jekyll clean

build: clean
	@${DEBUG} bundle exec jekyll build --profile --config _config.yml,.debug.yml

server: clean
	@${DEBUG} bundle exec jekyll server --livereload --config _config.yml,.debug.yml
```


3. _config.yml

```
title: Your project name
lang: en
description: a catchy description for your project

remote_theme: rundocs/jekyll-rtd-theme

readme_index:
  with_frontmatter: true

exclude:
  - Makefile
  - CNAME
  - Gemfile
  - Gemfile.lock
```

4. .debug.yml
```
remote_theme: false

theme: jekyll-rtd-theme
```