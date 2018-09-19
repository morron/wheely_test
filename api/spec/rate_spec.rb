require 'spec_helper'
require_relative '../lib/rate'

RSpec.describe Rate do
  subject { described_class.new(base: 250, min: 500) }

  it { expect(subject).to respond_to :base }
  it { expect(subject).to respond_to :min }
  it { expect(subject).to respond_to :perKm }
  it { expect(subject).to respond_to :perMin }

end
