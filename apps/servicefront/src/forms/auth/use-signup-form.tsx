import { addErrorsToForm } from "@/lib/utils";
import { AuthService, SignUpRequest } from "@/services/auth-service";
import { useMutation } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { FormEvent } from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";

const useSignUpForm = () => {
  const defaultValues = { email: "", password: "", first_name: "", last_name: "" };

  const form = useForm<SignUpRequest>({
    defaultValues,
    reValidateMode: "onBlur",
  });

  const signUpMutation = useMutation({
    mutationFn: AuthService.signUp,
    onSuccess: (data) => {
      if (data.success) {
        form.reset();
        toast.success("Account created successfully");
        // Add success logic here
      } else {
        if (data.errors) {
          addErrorsToForm(data.errors, form);
          // Add error logic here
        }
        toast.error(data.error);
        // Add error logic here
      }
    },
    onError: (error: AxiosError) => {},
  });

  const signUp = async (e: FormEvent) => {
    e.preventDefault();
    await signUpMutation.mutateAsync(form.getValues());
  };

  return [form, signUp] as const;
};

export default useSignUpForm;
