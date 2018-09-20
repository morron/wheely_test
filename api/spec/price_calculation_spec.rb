require 'spec_helper'
require_relative '../lib/service_pb'
require_relative '../lib/price_calculation'

RSpec.describe PriceCalculation do
  subject { described_class.new }
  let(:rate) { Rate.new(base: 250, perKm: 20, perMin: 20, min: 500) }
  let(:rate_with_min_price) { Rate.new(base: 250, perKm: 20, perMin: 20, min: 50_000) }
  let(:direction_result) { ::Wheelysvc::DirectionResult.new(distance: 25178, duration: 1957) }

  it "#call" do
    expect(subject.(rate, direction_result)).to eql 1406
  end

  it "returns min price" do
    expect(subject.(rate_with_min_price, direction_result)).to eql 50_000
  end
end
