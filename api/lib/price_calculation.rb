class PriceCalculation
  def call(ride, direction_result)
    distance_price = (direction_result.distance / 1000.0) * ride.perKm
    time_price = (direction_result.duration / 60.0) * ride.perMin

    price = (ride.base + distance_price + time_price).round

    price < ride.min ? ride.min : price
  end
end
