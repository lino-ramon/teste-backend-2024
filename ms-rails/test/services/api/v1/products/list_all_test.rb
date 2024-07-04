require 'test_helper'

class Services::Api::V1::Products::ListAllTest < ActiveSupport::TestCase
  def setup
    @params = {}
    @request = nil
    @service = Services::Api::V1::Products::ListAll.new(@params, @request)
  end

  test "should return all products" do
    products = Product.all
    assert_equal products, @service.execute
  end
end
