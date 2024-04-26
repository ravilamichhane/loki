import { FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input, InputProps } from "@/components/ui/input";
import { cn } from "@/lib/utils";
import { Control, FieldValues, Path } from "react-hook-form";

type TextareaInputProps<T extends FieldValues> = InputProps & {
  control: Control<T>;
  name: Path<T>;
  label?: string;
  wrapperClassnName?: HTMLDivElement["className"];
  description?: string;
};

export const TextInput = <T extends FieldValues>({
  control,
  name,
  label,
  description,
  wrapperClassnName,
  ...props
}: TextareaInputProps<T>) => {
  return (
    <FormField
      control={control}
      name={name}
      render={({ field }) => (
        <FormItem className={cn("grid", wrapperClassnName)}>
          {label && <FormLabel htmlFor={name}>{label}</FormLabel>}
          <FormControl>
            <Input
              {...props}
              id={name}
              value={field.value}
              onChange={(e) => {
                field.onChange(e);
                props.onChange?.(e);
              }}
            />
          </FormControl>

          {description && <FormDescription>{description}</FormDescription>}
          <FormMessage />
        </FormItem>
      )}
    />
  );
};
