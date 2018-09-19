require 'roda'
require 'json'

class App < Roda
  plugin :json_parser
  plugin :error_handler
  plugin :json, classes: [Array, Hash]

  route do |r|
    rate = Rate.new(base: 250, perKm: 20, perMin: 20, min: 500)

    r.root do
      "pong"
    end

    r.on 'calc' do
      result = ::ParamsSchema.call(r.params.to_h)

      if result.success?
        origin = ::Wheelysvc::Point.new(result.output['origin'])
        destination = ::Wheelysvc::Point.new(result.output['destination'])

        ride = ::Wheelysvc::Ride.new(origin: origin, destination: destination)

        grpc_response = GrpcClient.new.(ride)
        price = PriceCalculation.new.(rate, grpc_response)

        response = { price: price }
      end
      response = result.messages if result.failure?
      response
    rescue => err
      "Oh no! => #{err}"
    end
  end
end
