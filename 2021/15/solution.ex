defmodule Graph do
  def shortest_path(grid, source) do
    distances = %{}
    queue = MapSet.new(Enum.map(grid, fn {point, _weight} -> point end))
    distances = Map.update(distances, source, 0, fn _x -> 0 end)
    recur(distances, queue, grid)
  end

  defp recur(distances, queue, grid) do
    if MapSet.size(queue) == 0 do
      distances
    else
      {row, col} = u = Enum.min_by(queue, fn x -> Map.get(distances, x) end)
      queue = MapSet.delete(queue, u)

      neighbours = [
        {row, col + 1},
        {row, col - 1},
        {row + 1, col},
        {row - 1, col}
      ]
      |> Enum.filter(fn neighbour -> Enum.member?(queue, neighbour) end)

      distances = Enum.reduce(neighbours, distances, fn neighbour, distances ->
        dist_neighbour = Map.get(distances, neighbour)
        weight = grid[neighbour]
        new_distance = Map.get(distances, u) + weight
        if new_distance < dist_neighbour do
          Map.put(distances, neighbour, new_distance)
        else
          distances
        end
      end)

      recur(distances, queue, grid)
    end
  end
end

lines = File.read!("./input.txt")
  |> String.split("\n")

grid =
  for {line, row} <- Enum.with_index(lines),
      {weight, col} <- Enum.with_index(String.split(line, "", trim: true)),
      into: %{} do
    {{row, col}, String.to_integer(weight)}
  end

Graph.shortest_path(grid, {0, 0})
|> Enum.sort()
|> List.last()
|> IO.inspect()
