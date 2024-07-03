require 'test_helper'

module Services
  module Api
    module V1
      module Products
        class UpsertTest < ActiveSupport::TestCase
          def setup
            @params = {
              name: 'Test Product',
              brand: 'Test Brand',
              price: 99.99,
              description: 'Test Description',
              stock: 10
            }
            @is_api = false  # Assuming the is_api is not used in your service
          end

          test 'should create a new product' do
            service = Services::Api::V1::Products::Upsert.new(@params, @is_api)
            product = service.execute

            assert product.persisted?, 'Product should be saved to the database'
            assert_equal @params[:name], product.name
            assert_equal @params[:brand], product.brand
            assert_equal @params[:price], product.price
            assert_equal @params[:description], product.description
            assert_equal @params[:stock], product.stock
          end

          test 'should update an existing product' do
            existing_product = Product.create!(
              name: 'Old Name',
              brand: 'Old Brand',
              price: 49.99,
              description: 'Old Description',
              stock: 5
            )
            @params[:id] = existing_product.id
            @params[:name] = 'Updated Name'

            service = Services::Api::V1::Products::Upsert.new(@params, @is_api)
            product = service.execute

            assert product.persisted?, 'Product should be saved to the database'
            assert_equal 'Updated Name', product.name
            assert_equal @params[:brand], product.brand
            assert_equal @params[:price], product.price
            assert_equal @params[:description], product.description
            assert_equal @params[:stock], product.stock
          end
          
        end
      end
    end
  end
end
