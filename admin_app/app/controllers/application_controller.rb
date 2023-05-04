class ApplicationController < ActionController::Base
  def set_companies
    @companies = Company.all.order(name: :asc)
  end
end
