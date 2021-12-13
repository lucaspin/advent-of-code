defmodule Folder do
  def fold(dots, "x", value) do
    first_part =
      dots
      |> Enum.filter(fn [col, _row] -> col < value end)
      |> MapSet.new()

    second_part =
      dots
      |> Enum.filter(fn [col, _row] -> col > value end)
      |> Enum.map(fn [col, row] -> [col - value - 1, row] end)

    second_part_reversed =
      second_part
      |> Enum.map(fn [col, row] -> [abs(col - value + 1), row] end)
      |> MapSet.new()

    MapSet.union(first_part, second_part_reversed)
  end

  def fold(dots, "y", value) do
    first_part =
      dots
      |> Enum.filter(fn [_col, row] -> row < value end)
      |> MapSet.new()

    second_part =
      dots
      |> Enum.filter(fn [_col, row] -> row > value end)
      |> Enum.map(fn [col, row] -> [col, row - value - 1] end)

    second_part_reversed =
      second_part
      |> Enum.map(fn [col, row] -> [col, abs(row - value + 1)] end)
      |> MapSet.new()

    MapSet.union(first_part, second_part_reversed)
  end

  def print(dots) do
    Enum.each(0..max_row(dots), fn row ->
      Enum.each(0..max_col(dots), fn col ->
        if Enum.member?(dots, [col, row]) do
          IO.write("##")
        else
          IO.write("  ")
        end
      end)

      IO.write("\n")
    end)
  end

  def max_col(dots) do
    dots
      |> Enum.map(fn [col, _row] -> col end)
      |> Enum.max()
  end

  def max_row(dots) do
    dots
      |> Enum.map(fn [_col, row] -> row end)
      |> Enum.max()
  end
end

[dots, instructions] = File.read!("./input.txt")
  |> String.split("\n\n")
  |> Enum.map(fn line -> String.split(line, "\n") end)

visible_dots =
  dots
  |> Enum.map(fn dot ->
    dot
    |> String.split(",")
    |> Enum.map(&String.to_integer/1)
  end)

# Part 1
regex = ~r/fold along ([xy])=(\d+)$/
[first_instruction | _rest] = instructions
Enum.reduce([first_instruction], MapSet.new(visible_dots), fn instruction, visible_dots ->
  [_, direction, value] = Regex.run(regex, instruction)
  Folder.fold(visible_dots, direction, String.to_integer(value))
end)
|> Enum.count()
|> IO.inspect()

# Part 2
regex = ~r/fold along ([xy])=(\d+)$/
Enum.reduce(instructions, MapSet.new(visible_dots), fn instruction, visible_dots ->
  [_, direction, value] = Regex.run(regex, instruction)
  Folder.fold(visible_dots, direction, String.to_integer(value))
end)
|> Folder.print()
