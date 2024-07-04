class AddUuidToProducts < ActiveRecord::Migration[7.1]
  def change
    add_column :products, :uuid, :string, null: false, default: ""
    add_index :products, :uuid, unique: true
  end
end
