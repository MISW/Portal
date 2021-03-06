import { useEffect, useCallback, useMemo } from "react";
import { useSelector, useDispatch } from "react-redux";
import {
  selectPaymentPeriod,
  fetchPaymentPeriod,
  updatePaymentPeriod,
  selectCurrentPeriod,
  fetchCurrentPeriod,
  updateCurrentPeriod,
  selectEmailTemplateOf,
  fetchEmailTemplate,
  updateEmailTemplate as updateEmailTemplateThunk,
} from "features/appconfig";
import { EmailKind } from "models/appconfig";

export type PeriodConfigState = () => [
  number | undefined,
  (period: number) => Promise<void>
];

export const usePaymentPeriodConfig: PeriodConfigState = () => {
  const dispatch = useDispatch();
  const paymentPeriod = useSelector(selectPaymentPeriod);
  useEffect(() => {
    dispatch(fetchPaymentPeriod());
  }, [dispatch]);

  const update = useCallback(
    async (newPaymentPeriod: number) => {
      try {
        await dispatch(updatePaymentPeriod(newPaymentPeriod));
      } catch (e) {
        console.error(e);
        throw new Error("支払い期間の更新に失敗しました: " + e.message);
      }
    },
    [dispatch]
  );

  return [paymentPeriod, update];
};

export const useCurrentPeriodConfig: PeriodConfigState = () => {
  const dispatch = useDispatch();
  const currentPeriod = useSelector(selectCurrentPeriod);
  useEffect(() => {
    dispatch(fetchCurrentPeriod());
  }, [dispatch]);

  const update = useCallback(
    async (newCurrentPeriod: number) => {
      try {
        await dispatch(updateCurrentPeriod(newCurrentPeriod));
      } catch (e) {
        throw new Error("現在の期間の更新に失敗しました: " + e.message);
      }
    },
    [dispatch]
  );

  return [currentPeriod, update];
};

export type EmailTemplate = { subject: string; body: string };

export function useEmailTemplateConfig(emailKind: EmailKind) {
  const dispatch = useDispatch();
  const emailTemplate = useSelector(
    useMemo(() => selectEmailTemplateOf(emailKind), [emailKind])
  );
  useEffect(() => {
    if (emailTemplate == null) dispatch(fetchEmailTemplate(emailKind));
  }, [emailKind, emailTemplate, dispatch]);

  const updateEmailTemplate = useCallback(
    async (template: Readonly<EmailTemplate>) => {
      try {
        await dispatch(updateEmailTemplateThunk({ kind: emailKind, template }));
      } catch (e) {
        throw new Error(
          "Eメールテンプレートの更新に失敗しました: " + e.message
        );
      }
    },
    [emailKind, dispatch]
  );

  return {
    emailTemplate,
    updateEmailTemplate,
  } as const;
}
