FROM ruby:3.2

WORKDIR src
COPY . .

RUN bundle install

EXPOSE 3000

CMD ["./bin/rails", "server", "-b", "0.0.0.0"]

