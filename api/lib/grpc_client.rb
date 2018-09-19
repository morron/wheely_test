class GrpcClient
  def call(ride)
    stub = ::Wheelysvc::DistanceCalcultaion::Stub.new(ENV.fetch('GRPC_SERVER', 'localhost:50051'), :this_channel_is_insecure)

    stub.calculate(ride)
  rescue => err
    raise err
  end
end
