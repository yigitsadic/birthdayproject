class Employee < ApplicationRecord
  belongs_to :company

  validates :first_name, :last_name, :email, :birth_day, :birth_month, presence: true
  validates_uniqueness_of :email
end
