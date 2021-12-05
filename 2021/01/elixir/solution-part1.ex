defmodule Solution do
  def find do
    read_measurements() |> count(0)
  end

  def read_measurements do
    File.read!("./input.txt")
      |> String.split("\n")
      |> Enum.map(&String.to_integer/1)
  end

  def count([_ | []], acc), do: acc
  def count([previous | list], acc) do
    [current | _] = list
    if current > previous do
      count(list, acc + 1)
    else
      count(list, acc)
    end
  end
end

defmodule Solution2 do
  def run(input) do
    input
    |> String.split("\n")
    |> Enum.map(&String.to_integer/1)
    |> Enum.chunk_every(2, 1, :discard)
    |> Enum.count(fn [left, right] -> right > left end)
  end
end

IO.inspect(Solution2.run(File.read!("../input.txt")))
