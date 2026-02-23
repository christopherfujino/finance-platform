#!/usr/bin/env ruby

require 'csv'

local_path = ARGV[0]
if local_path.nil?
  STDERR.puts "Usage: csv-fixer.rb CSV_FILE"
  exit 1
end

parsed = []

# e.g. `["Aug 1, 2025", "Chase Sapphire Reserve", "Cko*patreon* Membership", "Fitness:Gym", "no", " -5.00"]`
CSV.foreach(local_path) do |row|
  a = blank = ''
  b = account = row[1]
  c = 'CLEARED'
  # Wrap in quotes, it has a comma
  d = date = "\"#{row[0]}\""
  e = payee = row[2]
  f = usage = ''
  g = category = row[3]
  h = tags = ''
  i = notes = ''
  j = amount = row[5]
  #exclusion = row[4]
  printf "#{a},#{b},#{c},#{d},#{e},#{f},#{g},#{h},#{i},#{j}\n"
end

# STDERR so we can pipe STDOUT to a file.
STDERR.puts "Done"
