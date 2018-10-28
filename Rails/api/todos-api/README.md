# README

* Fixed Model file constant class name missing error by running "spring stop", check out https://stackoverflow.com/questions/39438109/rails-console-in-production-nameerror-uninitialized-constant for details
* Run "rails runner 'puts Rails.env'" to check current rails env
* Run "rails console > ENV" to print out current ENVs
* Run "rails db:migrate -e production" to init production DB