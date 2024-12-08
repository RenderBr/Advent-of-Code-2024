# frozen_string_literal: true

class Test

  attr_accessor :test_value, :numbers

  def initialize
    @test_value = nil
    @numbers = []
    @memo = {}
  end

  def to_s
    "Test Value: #{@test_value}, Numbers: #{@numbers.join(', ')}"
  end

  def evaluate_test
    return false if @numbers.empty?

    numbers.each_with_index { |_, index|
      result = numbers[0]
      selected_nums = numbers[1...index + 1]

      return true if try(result,selected_nums)

    }

    false

  end

  def try(current, selected_nums)
    return false if selected_nums.empty?

    current_nums = selected_nums.dup

    while current_nums.any?
      current_num = current_nums.shift

      for op in [0,1]
        result = test_operators(current, current_num, op)

        return true if result == @test_value

        if try(result, current_nums)
          return true
        end


      end

    end

    false
  end

  def test_operators(current, other, op)
    case op
    when 0
      return current + other
    when 1
      return current * other
    end

    current
  end
end

