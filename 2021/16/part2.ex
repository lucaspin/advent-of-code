defmodule BITS do
  def evaluate do
    bits = bits()
    evaluate(bits)
  end

  def evaluate(bits) do
    {_version, rest} = version(bits)
    {type, packet} = type(rest)
    evaluate_type(type, 6, packet)
  end

  def evaluate_type(0, read, bits) do
    evalute_operation(bits, read, fn packet_values -> Enum.sum(packet_values) end)
  end

  def evaluate_type(1, read, bits) do
    evalute_operation(bits, read, fn packet_values -> Enum.product(packet_values) end)
  end

  def evaluate_type(2, read, bits) do
    evalute_operation(bits, read, fn packet_values -> Enum.min(packet_values) end)
  end

  def evaluate_type(3, read, bits) do
    evalute_operation(bits, read, fn packet_values -> Enum.max(packet_values) end)
  end

  def evaluate_type(4, read, bits) do
    find_literal_value("", read, bits)
  end

  def evaluate_type(5, read, bits) do
    evalute_operation(bits, read, fn [value2, value1] ->
      if value1 > value2, do: 1, else: 0
    end)
  end

  def evaluate_type(6, read, bits) do
    evalute_operation(bits, read, fn [value2, value1] ->
      if value1 < value2, do: 1, else: 0
    end)
  end

  def evaluate_type(7, read, bits) do
    evalute_operation(bits, read, fn [value2, value1] ->
      if value1 == value2, do: 1, else: 0
    end)
  end

  defp find_literal_value(final, read, <<1::1, value::4, rest::bitstring>>) do
    find_literal_value(
      final <> String.pad_leading(Integer.to_string(value, 2), 4, "0"),
      read + 5,
      rest
    )
  end

  defp find_literal_value(final, read, <<0::1, value::4, rest::bitstring>>) do
    final = final <> String.pad_leading(Integer.to_string(value, 2), 4, "0")
    {String.to_integer(final, 2), rest, read + 5}
  end

  defp evalute_operation(<<0::1, rest::bitstring>>, read, fun) do
    <<bit_length::15, sub_packets::bitstring>> = rest
    {packet_values, rest} = evaluate_bits_until([], sub_packets, bit_length)
    {fun.(packet_values), rest, read + 16 + bit_length}
  end

  defp evalute_operation(<<1::1, rest::bitstring>>, read, fun) do
    <<num_packets::11, sub_packets::bitstring>> = rest
    {packet_values, rest, bits_read} = evaluate_packets([], sub_packets, num_packets, read + 12)
    {fun.(packet_values), rest, bits_read}
  end

  defp evaluate_packets(packet_values, bits, 0, read), do: {packet_values, bits, read}
  defp evaluate_packets(packet_values, bits, packet_count, read) do
    {packet_value, rest, bits_read} = evaluate(bits)
    evaluate_packets([packet_value | packet_values], rest, packet_count - 1, bits_read + read)
  end

  defp evaluate_bits_until(packet_values, bits, read) when read <= 0, do: {packet_values, bits}
  defp evaluate_bits_until(packet_values, bits, read) do
    {packet_value, rest, bits_read} = evaluate(bits)
    evaluate_bits_until([packet_value | packet_values], rest, read - bits_read)
  end

  def version(bits) do
    <<version::3, rest::bitstring>> = bits
    {version, rest}
  end

  def type(bits) do
    <<type::3, rest::bitstring>> = bits
    {type, rest}
  end

  def bits do
    File.read!("./input.txt")
    |> Base.decode16!()
  end
end

BITS.evaluate()
|> elem(0)
|> IO.inspect()
