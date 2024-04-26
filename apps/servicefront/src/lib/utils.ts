import { clsx, type ClassValue } from "clsx";
import { FieldValues, useForm } from "react-hook-form";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}
// TODO : FIX TYPE SAFETY
export function addErrorsToForm<T extends FieldValues>(
  errors: Record<string, string>,
  form: ReturnType<typeof useForm<T>>,
) {
  form.clearErrors();
  Object.entries(errors).forEach(([key, value]) => {
    form.setError(key as any, { message: value });
  });
}
