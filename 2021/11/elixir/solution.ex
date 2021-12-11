defmodule Solution do
  def recur(grid, points_to_flash) do
    recur(grid, points_to_flash, [], 0)
  end

  def recur(grid, [], _, flash_count), do: {grid, flash_count}
  def recur(grid, [{row, col} | rest], visited, flash_count) do
    neighbours = [
      {row - 1, col},
      {row + 1, col},
      {row, col - 1},
      {row, col + 1},
      {row - 1, col - 1},
      {row - 1, col + 1},
      {row + 1, col - 1},
      {row + 1, col + 1}
    ]

    grid = grid
      |> Enum.map(fn {point, value} ->
        if Enum.member?(neighbours, point) do
          {point, value + 1}
        else
          {point, value}
        end
      end)
      |> Enum.into(%{})

    new_rest = grid
    |> Enum.filter(fn {point, _value} -> Enum.member?(neighbours, point) end)
    |> Enum.filter(fn {_point, value} -> value > 9 end)
    |> Enum.filter(fn {point, _value} -> !Enum.member?(visited, point) end)
    |> Enum.map(fn {point, _value} -> point end)
    |> Enum.reduce(MapSet.new(rest), fn point, set -> MapSet.put(set, point) end)
    |> MapSet.to_list()

    recur(grid, new_rest, [{row,col} | visited], flash_count + 1)
  end

  def all_flashed?(grid) do
    grid
    |> Enum.map(fn {_point, value} -> value end)
    |> Enum.all?(fn value -> value == 0 end)
  end

  def increase_all(grid) do
    grid
    |> Enum.map(fn {k, v} -> {k, v+1} end)
    |> Enum.into(%{})
  end

  def find_above(grid, value) do
    grid
    |> Enum.filter(fn {_k, v} -> v > value end)
    |> Enum.map(fn {point, _value} -> point end)
  end
end

lines = File.read!("../input.txt")
  |> String.split("\n")
  |> Enum.map(fn line ->
    line
    |> String.split("", trim: true)
    |> Enum.map(&String.to_integer/1)
  end)

grid =
  for {line, row} <- Enum.with_index(lines),
      {number, col} <- Enum.with_index(line),
      into: %{} do
    {{row, col}, number}
  end

# Part 1
Enum.reduce(1..100, %{grid: grid, flashes: 0}, fn _step_number, %{grid: grid, flashes: flashes} ->
  increased_grid = Solution.increase_all(grid)
  points_to_flash = Solution.find_above(increased_grid, 9)
  {flashed_grid, count} =
    Solution.recur(increased_grid, points_to_flash)

  final_grid =
    flashed_grid
    |> Enum.map(fn {k, v} -> if v > 9, do: {k, 0}, else: {k, v} end)
    |> Enum.into(%{})

  %{grid: final_grid, flashes: flashes + count}
end)
|> Map.get(:flashes)
|> IO.inspect()

# Part 2
Enum.reduce_while(1..1000, %{grid: grid, step: 0}, fn _step_number, %{grid: grid, step: step} ->
  increased_grid = Solution.increase_all(grid)
  points_to_flash = Solution.find_above(increased_grid, 9)
  {flashed_grid, _count} =
    Solution.recur(increased_grid, points_to_flash)

  final_grid =
    flashed_grid
    |> Enum.map(fn {k, v} -> if v > 9, do: {k, 0}, else: {k, v} end)
    |> Enum.into(%{})

  if Solution.all_flashed?(grid) do
    {:halt, %{grid: final_grid, step: step}}
  else
    {:cont, %{grid: final_grid, step: step + 1}}
  end
end)
|> Map.get(:step)
|> IO.inspect()
