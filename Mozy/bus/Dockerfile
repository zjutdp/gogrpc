FROM ubuntu:14.04
RUN useradd -ms /bin/bash bus && mkdir -p /usr/local/etc && {            echo 'install: --no-document';          echo 'update: --no-document';   } >> /usr/local/etc/gemrc
RUN apt-get update && apt-get install -y ruby 1.9.3 && apt-get install git g++ gcc make libc6-dev libreadline6-dev zlib1g-dev libssl-dev libyaml-dev libsqlite3-dev sqlite3 autoconf libgmp-dev libgdbm-dev libncurses5-dev automake libtool bison pkg-config libffi-dev libxml2-dev libpq-dev libmagickwand-dev -y --force-yes
RUN gem install bundler

USER bus
WORKDIR /home/bus
RUN pwd
COPY . $WORKDIR

USER root
RUN bundle install

WORKDIR vendor/gems/rdiscount-1.6.8/ext/
RUN ruby extconf.rb
RUN make
#RUN cp rdiscount.bundle ../lib/
RUN cp rdiscount.so ../lib/
WORKDIR /home/bus
RUN cp vendor/plugins/payjunction/config/merchant_accounts.yml.dev vendor/plugins/payjunction/config/merchant_accounts.yml

#USER bus
WORKDIR /home/bus
