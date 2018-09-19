require_relative 'app'

this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)
Dir[File.join(File.dirname(__FILE__), 'lib', '**', '*.rb')].sort.each {|file| require file }

run App.freeze.app
