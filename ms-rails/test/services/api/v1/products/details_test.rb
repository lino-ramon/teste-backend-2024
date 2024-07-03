require 'test_helper'

class Services::Api::V1::Products::DetailsTest < ActiveSupport::TestCase
  def setup
    @product = products(:one)
  end

  test "should raise error if product is not found" do
    params = { id: '9999' }
    service = Services::Api::V1::Products::Details.new(params, nil)
    assert_raises(ActiveRecord::RecordNotFound) { service.execute }
  end

  test "should return product if found" do
    params = { id: @product.id }
    service = Services::Api::V1::Products::Details.new(params, nil)
    result = service.execute
    assert_equal @product, result
  end
end
