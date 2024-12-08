# frozen_string_literal: true

require_relative 'test'

def process_tests(filename)
  tests = []

  File.foreach(filename) do |line|
    current_test = Test.new

    left = ""
    right = ""
    passed_test_value = false

    line.chars.each { |char|
      if char == ':'
        passed_test_value = true
        current_test.test_value = left.to_i
        next
      end

      unless passed_test_value
        left += char
        next
      end

      if passed_test_value
        if char == ' '
          if right != ""
            current_test.numbers.push(right.to_i)
          end

          right = ""
          next
        end

        right += char
      end
    }

    current_test.numbers.push(right.to_i) unless right.empty?

    tests.push(current_test)
  end

  tests
end

tests = process_tests("data.txt")

sum = 0

for test in tests
  if test.evaluate_test
    sum += test.test_value
  end
end

puts sum


