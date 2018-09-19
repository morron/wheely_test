class Rate
  attr_reader :base, :min, :perKm, :perMin

  def initialize(**opts)
    @base = opts[:base]
    @min = opts[:min]
    @perMin = opts[:perMin]
    @perKm = opts[:perKm]
  end
end
