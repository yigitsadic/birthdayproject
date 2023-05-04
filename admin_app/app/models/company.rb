class Company < ApplicationRecord
  has_many :users, dependent: :delete_all
  has_many :employees, dependent: :delete_all
end
