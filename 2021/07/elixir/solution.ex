defmodule FuelCalculator do
  def part_one do
    reduce(fn n, m -> abs(m - n) end)
  end

  def part_two do
    reduce(fn n, m -> 0..abs(m - n) |> Enum.sum() end)
  end

  defp reduce(fun) do
    numbers = read()
    max = numbers |> Enum.reduce(0, fn n, max -> if n > max, do: n, else: max end)

    Enum.reduce(0..max, %{}, fn n, state ->
      if Map.has_key?(state, n) do
        state
      else
        Map.put_new(state, n, Enum.map(numbers, fn m -> fun.(n, m) end))
      end
    end)
    |> Enum.map(fn {_k, v} -> Enum.sum(v) end)
    |> Enum.min()
  end

  defp read do
    File.read!("../input.txt")
      |> String.split(",")
      |> Enum.map(&String.to_integer/1)
  end
end

IO.inspect(FuelCalculator.part_one())
IO.inspect(FuelCalculator.part_two())
