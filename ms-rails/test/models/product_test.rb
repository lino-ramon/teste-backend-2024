require 'test_helper'

class ProductTest < ActiveSupport::TestCase
  def setup
    @product = Product.new(
      name: 'Test Product',
      price: 99.99,
      brand: 'Test Brand',
      description: 'Test Description',
      stock: 10
    )
  end

  test "should be valid with valid attributes" do
    assert @product.valid?
  end

  test "name should be present and have at least 4 characters" do
    @product.name = "abc"
    assert_not @product.valid?

    @product.name = "abcd"
    assert @product.valid?
  end

  test "price should be present, a valid format, and between 0 and 1,000,000" do
    @product.price = nil
    assert_not @product.valid?

    @product.price = -1
    assert_not @product.valid?

    @product.price = 1_000_000
    assert_not @product.valid?

    @product.price = 9.9999
    assert_not @product.valid?

    @product.price = 9.99
    assert @product.valid?
  end

  test "brand should be present" do
    @product.brand = ""
    assert_not @product.valid?
  end

  test "description should be present" do
    @product.description = ""
    assert_not @product.valid?
  end

  test "stock should be an integer and at least 0" do
    @product.stock = -1
    assert_not @product.valid?

    @product.stock = 0
    assert @product.valid?

    @product.stock = 10
    assert @product.valid?
  end
end
