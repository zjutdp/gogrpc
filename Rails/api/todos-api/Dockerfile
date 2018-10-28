FROM ruby:2.3

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app
COPY Gemfile* ./
RUN bundle install
COPY . .
ENV SECRET_KEY_BASE="7680cc63beaf8a27b735597664a2251105dc011af977d79565443ef147c2d84840618bb8c42b846692387502647c82b0951d0ee343686cd69ffc23ac2938b5cc"
EXPOSE 3000
CMD ["rails", "server", "-e", "production",  "-b", "0.0.0.0"]
