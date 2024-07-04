class RemoveUuidFromProducts < ActiveRecord::Migration[7.1]
  def change
    remove_column :products, :uuid, :string
  end
end
