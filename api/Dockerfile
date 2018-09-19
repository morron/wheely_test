FROM ruby:2.5.1

WORKDIR /app

ENV BUNDLE_PATH=/bundle \
    BUNDLE_BIN=/bundle/bin \
    GEM_HOME=/bundle

ENV PATH="${BUNDLE_BIN}:${PATH}"

COPY Gemfile* /app/
RUN bundle install
COPY . /app

EXPOSE 9393
CMD ["shotgun", "-o", "0.0.0.0"]

