# README
* Great article @ https://scotch.io/tutorials/build-a-restful-json-api-with-rails-5-part-one#toc-conclusion
* Fixed Model file constant class name missing error by running "spring stop", check out https://stackoverflow.com/questions/39438109/rails-console-in-production-nameerror-uninitialized-constant for details
* Fixed SECRET_KEY_BASE missing error by https://stackoverflow.com/questions/23180650/how-to-solve-error-missing-secret-key-base-for-production-environment-rai
* Run "rails runner 'puts Rails.env'" to check current rails env
* Run "rails console > ENV" to print out current ENVs
* Run "rails db:migrate -e production" to init production DB
* Run "bin/rails r 'puts ActiveSupport::Dependencies.autoload_paths'" to print auto loading paths
* Use "ab" cmd http://httpd.apache.org/docs/2.4/programs/ab.html for perf testing
