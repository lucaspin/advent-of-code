defmodule Solution do
  def run(input) do
    input
    |> String.split("\n")
    |> Enum.reduce({_h_pos=0, _v_pos=0, _aim=0}, fn
      "forward " <> number, {h_pos, v_pos, aim} ->
        {h_pos+String.to_integer(number), v_pos+(aim*String.to_integer(number)), aim}
      "up " <> number, {h_pos, v_pos, aim} ->
        {h_pos, v_pos, aim-String.to_integer(number)}
      "down " <> number, {h_pos, v_pos, aim} ->
        {h_pos, v_pos, aim+String.to_integer(number)}
    end)
    |> then(fn {h_pos, v_pos, _aim} -> h_pos * v_pos end)
  end
end

IO.inspect(Solution.run(File.read!("../input.txt")))
