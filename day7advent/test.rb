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

  # example: 4900956092109: 699 8 3 36 5 4 4 458 7 7
  def evaluate_test
    return false if @numbers.empty?

    starting_num = numbers[0]
    selected_nums = numbers[1..]
    return true if try(starting_num, selected_nums)

    false

  end

  def try(current, selected_nums)
    return false if selected_nums.empty?

    current_num = selected_nums.first
    remaining_nums = selected_nums[1..-1]

    #addition
    result_add = current + current_num
    return true if result_add == @test_value && remaining_nums.length == 0

    try_add_path = try(result_add, remaining_nums)
    return try_add_path if try_add_path

    #mult
    result_mult = current*current_num
    return true if result_mult == @test_value && remaining_nums.length == 0

    try_mult_path = try(result_mult, remaining_nums)
    return try_mult_path if try_mult_path

    #append
    result_append = "#{current}#{current_num}".to_i
    return true if result_append == @test_value && remaining_nums.length == 0

    try_append_path = try(result_append, remaining_nums)
    return try_append_path if try_append_path

    false
  end

  def english_op(op)
    op == 0 ? "+" : "*"
  end
end

