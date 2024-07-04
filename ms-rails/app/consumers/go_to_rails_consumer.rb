# frozen_string_literal: true

class GoToRailsConsumer < ApplicationConsumer
  def consume
    Karafka.logger.info "GoToRails - Consumer running"
    messages.each do |message|
      params = message.payload.with_indifferent_access

      begin
        upsert_service = Services::Api::V1::Products::Upsert.new(params, false)
        upsert_service.execute

        Karafka.logger.info "Successfully processed message with payload: #{params}"
      rescue => e
        Karafka.logger.error "Failed to process message with payload: #{params}, error: #{e.message}"
      end
    end
  end
end

