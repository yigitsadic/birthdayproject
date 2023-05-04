Rails.application.routes.draw do
  resources :companies
  resources :employees
  
  root "dashboard#index"
end
