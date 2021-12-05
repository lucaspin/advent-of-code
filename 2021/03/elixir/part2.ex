defmodule Solution do
  def o2(numbers) do
    recursion(numbers, 0, fn zero_count, one_count ->
      if zero_count > one_count, do: ?0, else: ?1
    end)
  end

  def co2(numbers) do
    recursion(numbers, 0, fn zero_count, one_count ->
      if one_count >= zero_count, do: ?0, else: ?1
    end)
  end

  defp recursion([number], _, _) do
    number
    |> Tuple.to_list()
    |> List.to_integer(2)
  end

  defp recursion(numbers, index, fun) do
    total = length(numbers)
    zero_count = Enum.count(numbers, fn n -> elem(n, index) == ?0 end)
    one_count = total - zero_count
    bit = fun.(zero_count, one_count)
    numbers = Enum.filter(numbers, fn n -> elem(n, index) == bit end)
    recursion(numbers, index + 1, fun)
  end
end

numbers = File.read!("../input.txt")
  |> String.split("\n")
  |> Enum.map(fn e -> String.to_charlist(e) end)
  |> Enum.map(fn e -> List.to_tuple(e) end)

IO.puts(Solution.o2(numbers) * Solution.co2(numbers))
