class ApplicationController < ActionController::Base
  http_basic_authenticate_with name: ENV['ADMIN_USER_NAME'], password: ENV['ADMIN_PASSWORD']

  def set_companies
    @companies = Company.all.order(name: :asc)
  end
end
