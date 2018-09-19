require 'dry-validation'

PointSchema = Dry::Validation.Schema do
  configure { config.type_specs = true }

  required('lat', :float).filled(:float?)
  required('lng', :float).filled(:float?)
end

ParamsSchema = Dry::Validation.Schema do
  required('origin').schema(PointSchema)
  required('destination').schema(PointSchema)
end
