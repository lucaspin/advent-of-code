defmodule GridCounter do
  def ignoring_diagonals do
    read()
    |> Enum.reduce(%{}, fn
      [x, y1, x, y2], grid -> update_map(Stream.cycle([x]), y1..y2, grid)
      [x1, y, x2, y], grid -> update_map(x1..x2, Stream.cycle([y]), grid)
      _, grid -> grid
    end)
    |> Enum.count(fn {_, value} -> value > 1 end)
  end

  def with_diagonals do
    read()
    |> Enum.reduce(%{}, fn
      [x, y1, x, y2], grid -> update_map(Stream.cycle([x]), y1..y2, grid)
      [x1, y, x2, y], grid -> update_map(x1..x2, Stream.cycle([y]), grid)
      [x1, y1, x2, y2], grid -> update_map(x1..x2, y1..y2, grid)
    end)
    |> Enum.count(fn {_, value} -> value > 1 end)
  end

  defp update_map(xs, ys, grid) do
    Enum.reduce(Enum.zip(xs, ys), grid, fn {x, y}, grid ->
      Map.update(grid, {x, y}, 1, &(&1 + 1))
    end)
  end

  defp read do
    File.read!("../input.txt")
      |> String.split("\n", trim: true)
      |> Enum.map(fn line -> String.split(line, [" -> ", ","], trim: true) end)
      |> Enum.map(fn row -> Enum.map(row, &String.to_integer/1) end)
  end
end

IO.inspect(GridCounter.ignoring_diagonals())
IO.inspect(GridCounter.with_diagonals())
