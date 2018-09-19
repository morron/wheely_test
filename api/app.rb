require 'roda'
require 'json'

class App < Roda
  plugin :json_parser
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
          stub = ::Wheelysvc::DistanceCalcultaion::Stub.new(ENV.fetch('GRPC_SERVER', 'localhost:50051'), :this_channel_is_insecure)

          resp = stub.calculate(ride)

          price = PriceCalculation.new.(rate, resp)

          response = { price: price }
        end
        response = result.messages if result.failure?
        response
    end
  end
end
