defmodule Solution do
  def run(input) do
    numbers = input
      |> String.split("\n")
      |> Enum.map(fn e -> String.to_charlist(e) end)
      |> Enum.map(fn e -> List.to_tuple(e) end)

    [sample | _] = numbers
    num_digits = tuple_size(sample)
    total_lines = length(numbers)

    gamma_as_list = for column <- 0..num_digits-1 do
      one_count = Enum.count(numbers, fn number -> elem(number, column) == ?1 end)
      zero_count = total_lines - one_count
      if zero_count > one_count, do: ?0, else: ?1
    end

    epsilon_as_list = gamma_as_list
      |> Enum.map(fn c -> if c == ?1, do: ?0, else: ?1 end)

    gamma = List.to_integer(gamma_as_list, 2)
    epsilon = List.to_integer(epsilon_as_list, 2)
    IO.puts("Gamma: #{gamma}, epsilon: #{epsilon}, result: #{gamma * epsilon}")
  end
end

Solution.run(File.read!("../input.txt"))
