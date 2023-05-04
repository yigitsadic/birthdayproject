Rails.application.routes.draw do
  resources :users
  resources :companies
  resources :employees
  
  root "dashboard#index"
end
